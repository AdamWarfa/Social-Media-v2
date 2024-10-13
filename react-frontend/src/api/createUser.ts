import { AuthorType } from "../models/author";

const endpoint = "http://localhost:4000";

async function postUser(userObject: AuthorType) {
  const json = JSON.stringify(userObject);

  console.log(json);

  const response = await fetch(`${endpoint}/users/register`, {
    method: "POST",
    body: json,
    headers: {
      "Content-Type": "application/json",
    },
  });

  return response.json();
}

export { postUser };
