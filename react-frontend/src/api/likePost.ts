import { PostType } from "../models/post";

const endpoint = "http://localhost:4000";

async function likePost(post: PostType) {
  const token = localStorage.getItem("token");

  const response = await fetch(`${endpoint}/api/posts/${post.id}/like`, {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
      Authorization: `Bearer ${token}`,
    },
  });
  if (response.ok) {
    return response.json();
  } else {
    console.log("error");
  }
}

async function unlikePost(post: PostType) {
  const token = localStorage.getItem("token");
  const response = await fetch(`${endpoint}/api/posts/${post.id}/unlike`, {
    method: "DELETE",
    headers: {
      "Content-Type": "application/json",
      Authorization: `Bearer ${token}`,
    },
  });
  if (response.ok) {
    return response.json();
  } else {
    console.log("error");
  }
}

export { likePost, unlikePost };
