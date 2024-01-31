const endpoint = "http://localhost:4000";

async function deletePost(postId: string) {
  const response = await fetch(`${endpoint}/posts/${postId}`, {
    method: "DELETE",
  });
  return response;
}

export {deletePost};