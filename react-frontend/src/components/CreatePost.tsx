import { useState, FormEvent } from "react";
import { uploadPost } from "../api/createPost";
import { PostRequest } from "../models/post";
import { useAuth } from "../security/AuthProvider";

interface CreatePostProps {
  loggedIn: boolean;
  userId: string;
}

export default function CreatePost({ userId }: CreatePostProps) {
  const [isDialogOpen, setIsDialogOpen] = useState(false);
  const auth = useAuth();

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
    const newPost: PostRequest = {
      text: (form.caption as HTMLTextAreaElement).value,
      imgSrc: (form.imgUrl as HTMLInputElement).value,
      authorId: userId,
    };

    uploadPost(newPost);
    console.log(newPost);
    handleCloseDialog();
  };

  return (
    <>
      {auth.isLoggedIn() && (
        <>
          <div className="flex justify-center items-center m-5">
            <div className="w-20 h-20 p-16 bg-black-100 transition-transform rounded-xl flex justify-center items-center" onClick={handleOpenDialog}>
              <h2 className="p-0 text-8xl text-black-950">+</h2>
            </div>
          </div>
        </>
      )}

      {isDialogOpen && (
        <>
          <section className="fixed top-0 left-0 z-50 w-full h-full bg-gray-900 bg-opacity-50 flex justify-center items-center">
            <div className=" bg-white editor mx-auto w-10/12 flex flex-col text-gray-800 rounded-lg p-4 shadow-lg max-w-2xl">
              <div className="heading text-center font-bold text-2xl m-5 text-gray-800">New Post</div>
              <form action="" method="post" onSubmit={handleCreatePost}>
                <div className="flex flex-col mt-4">
                  <input type="url" name="imgUrl" id="imgUrl" placeholder="Image URL" className="title bg-gray-200 border border-gray-300 p-2 mb-4 outline-none" spellCheck="false" />
                </div>
                <div className="flex flex-col mt-4">
                  <textarea name="caption" id="caption" placeholder="Caption" className="description bg-gray-200 sec p-3 h-60 border border-gray-300 outline-none" spellCheck="false" />
                </div>
                <div className="icons flex text-gray-500 m-2">
                  <svg className="mr-2 cursor-pointer hover:text-gray-700 border rounded-full p-1 h-7" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17.657 16.657L13.414 20.9a1.998 1.998 0 01-2.827 0l-4.244-4.243a8 8 0 1111.314 0z" />
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 11a3 3 0 11-6 0 3 3 0 016 0z" />
                  </svg>
                  <svg className="mr-2 cursor-pointer hover:text-gray-700 border rounded-full p-1 h-7" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M14.828 14.828a4 4 0 01-5.656 0M9 10h.01M15 10h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z" />
                  </svg>
                  <svg className="mr-2 cursor-pointer hover:text-gray-700 border rounded-full p-1 h-7" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15.172 7l-6.586 6.586a2 2 0 102.828 2.828l6.414-6.586a4 4 0 00-5.656-5.656l-6.415 6.585a6 6 0 108.486 8.486L20.5 13" />
                  </svg>
                  <div className="count ml-auto text-gray-400 text-xs font-semibold">{}/300</div>
                </div>
                <div className="buttons flex">
                  <button type="button" onClick={handleCloseDialog} className="btn border border-gray-300 p-1 px-4 font-semibold cursor-pointer text-gray-500 ml-auto">
                    Cancel
                  </button>
                  <button className="btn border border-indigo-500 p-1 px-4 font-semibold cursor-pointer text-gray-200 ml-2 bg-indigo-500">Post</button>
                </div>
              </form>
            </div>
          </section>
        </>
      )}

      {/* {isDialogOpen && (
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
      )} */}
    </>
  );
}
