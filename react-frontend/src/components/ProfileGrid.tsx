import { useState, useEffect } from "react";
import { getPostsByAuthor } from "../api/get.ts";
import PostType from "../models/post.ts";
import Post from "./Post";

interface ProfileGridProps {
  authorId?: string;
}

export default function ProfileGrid({ authorId }: ProfileGridProps) {
  const [postList, setPostList] = useState<PostType[]>([]);
  useEffect(() => {
    function createPostList() {
      getPostsByAuthor(authorId).then((data) => {
        setPostList(data);
      });
    }
    createPostList();
  }, []);

  return (
    <div className="py-12 flex items-center justify-center">
      <div className="grid grid-cols-1 md:grid-cols-2 xl:grid-cols-3 gap-2">
        {postList.map((post) => (
          <Post key={post.id} post={post} />
        ))}
      </div>
    </div>
  );
}
