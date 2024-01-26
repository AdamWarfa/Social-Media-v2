import { useState, FormEvent } from "react";
import { uploadPost } from "../api/createPost";

interface CreatePostProps {
  loggedIn: boolean;
  userId: string;
}

export default function CreatePost({ loggedIn, userId }: CreatePostProps) {
  const [isDialogOpen, setIsDialogOpen] = useState(false);

  const handleOpenDialog = () => {
    setIsDialogOpen(true);
  };

  const handleCloseDialog = () => {
    setIsDialogOpen(false);
  };

  const handleCreatePost = (event: FormEvent<HTMLFormElement>) => {
    event.preventDefault(); // Prevent default form submission behavior

    // Access form elements
    const form = event.currentTarget;
    // Create post
    const newPost = {
      id: Math.random().toString(36),
      text: (form.caption as HTMLTextAreaElement).value,
      imgSrc: (form.imgUrl as HTMLInputElement).value,
      author: userId,
      likes: 0,
      postDate: new Date().toString(),
    };

    uploadPost(newPost);
    console.log(newPost);
  };

  return (
    <>
      {loggedIn && (
        <>
          <h2>Create Post</h2>
          <div className="flex justify-center items-center m-5">
            <div className="w-20 h-20 p-20 gradient-button transition-transform rounded-full flex justify-center items-center" onClick={handleOpenDialog}>
              <h2 className="p-0 text-8xl text-slate-800">+</h2>
            </div>
          </div>
        </>
      )}

      {isDialogOpen && (
        <section className="fixed top-0 left-0 w-full h-full bg-gray-900 bg-opacity-50 flex justify-center items-center">
          <div className="bg-white rounded-lg w-1/2">
            <form action="" method="post" onSubmit={handleCreatePost}>
              <div className="flex flex-col p-8">
                <h2 className="text-3xl text-gray-900 font-bold text-center">Create Post</h2>
                <div className="flex flex-col mt-4">
                  <label htmlFor="imgUrl" className="text-lg text-gray-900 font-semibold mb-2">
                    Image URL
                  </label>
                  <input type="url" name="imgUrl" id="imgUrl" className="border-2 border-gray-200 p-2 rounded-lg focus:outline-none focus:border-slate-800" />
                </div>
                <div className="flex flex-col mt-4">
                  <label htmlFor="caption" className="text-lg text-gray-900 font-semibold mb-2">
                    Caption
                  </label>
                  <textarea name="caption" id="caption" className="border-2 border-gray-200 p-2 rounded-lg focus:outline-none focus:border-slate-800" />
                </div>
                <div className="flex justify-center mt-8">
                  <button className="gradient-button w-1/2 p-3 rounded-lg">Create Post</button>
                </div>
                <div className="flex justify-center mt-8">
                  <button className="gradient-button w-1/2 p-3 rounded-lg" type="button" onClick={handleCloseDialog}>
                    Close
                  </button>
                </div>
              </div>
            </form>
          </div>
        </section>
      )}
    </>
  );
}
