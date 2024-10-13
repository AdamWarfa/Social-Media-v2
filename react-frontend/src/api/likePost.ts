import { PostType } from "../models/post";

const endpoint = "http://localhost:4000";

async function likePost(post: PostType) {
  const json = JSON.stringify(post);

  const response = await fetch(`${endpoint}/posts/${post.id}`, {
    method: "PUT",
    body: json,
    headers: {
      "Content-Type": "application/json",
    },
  });
  if (response.ok) {
    return response.json();
  } else {
    console.log("error");
  }
}

export { likePost };
