import PostType from "../models/post";
const endpoint = "http://localhost:4000";

async function uploadPost(postObject: PostType) {
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
