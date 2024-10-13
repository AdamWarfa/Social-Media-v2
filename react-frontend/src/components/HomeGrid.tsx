import { useState, useEffect } from "react";
import { getPosts } from "../api/get.ts";
import { PostType } from "../models/post.ts";
import Post from "./Post";
import CreatePost from "./CreatePost.tsx";

interface HomeGridProps {
  userId: string;
  loggedIn: boolean;
}

export default function HomeGrid({ userId, loggedIn }: HomeGridProps) {
  const [postList, setPostList] = useState<PostType[] | undefined>(undefined); // Initialize as undefined
  const [sortOption, setSortOption] = useState("new");

  useEffect(() => {
    // Fetch posts when the component mounts or when sortOption changes
    fetchPosts();
  }, [sortOption]);

  const fetchPosts = async () => {
    try {
      const posts = await getPosts(sortOption);
      setPostList(posts);
    } catch (error) {
      console.error("Error fetching posts:", error);
    }
  };

  const handleSortChange = (e: React.ChangeEvent<HTMLSelectElement>) => {
    setSortOption(e.target.value);
  };

  return (
    <>
      <div className="py-12 flex items-center justify-center">
        <div className="mt-20 grid grid-cols-1 md:grid-cols-1 xl:grid-cols-1 gap-2">
          <CreatePost loggedIn={loggedIn} userId={userId} />
          <select className="rounded-lg my-2" name="sortHome" id="sortHome" value={sortOption} onChange={handleSortChange}>
            <option value="new">New</option>
            <option value="popular">Popular</option>
          </select>

          {postList && postList.map((post) => <Post key={post.id} post={post} />)}
        </div>
      </div>
    </>
  );
}
