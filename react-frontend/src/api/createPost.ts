import { PostRequest } from "../models/post";
const endpoint = "http://localhost:4000";

async function uploadPost(postObject: PostRequest) {
  const json = JSON.stringify(postObject);

  const response = await fetch(`${endpoint}/posts`, {
    method: "POST",
    body: json,
    headers: {
      "Content-Type": "application/json",
    },
  });
  return response;
}

export { uploadPost };
