import { PostType } from "../models/post";

const endpoint: string = "http://localhost:4000";

async function getPosts(sortOption: string) {
  const res = await fetch(`${endpoint}/posts`);
  const data = await res.json();

  switch (sortOption) {
    case "popular":
      data.sort((a: PostType, b: PostType) => b.likes - a.likes);
      break;
    case "new":
      data.sort((a: PostType, b: PostType) => new Date(b.postDate).getTime() - new Date(a.postDate).getTime());
      break;
  }
  return data;
}

async function getAuthor(id: string) {
  const res = await fetch(`${endpoint}/users/${id}`);
  const data = await res.json();
  return data;
}

async function getPostsByAuthor(id: string | undefined) {
  const res = await fetch(`${endpoint}/posts/author/${id}`);
  const data = await res.json();
  return data;
}

async function getLikeCount(post: PostType) {
  const res = await fetch(`${endpoint}/posts/${post.id}/like/count`);
  const data = await res.json();
  return data;
}

async function getHasLiked(post: PostType) {
  const token = localStorage.getItem("token");
  const res = await fetch(`${endpoint}/api/posts/${post.id}/hasliked`, {
    headers: {
      Authorization: `Bearer ${token}`,
    },
  });
  const data = await res.json();
  return data;
}

export { getPosts, getAuthor, getPostsByAuthor, getLikeCount, getHasLiked };
