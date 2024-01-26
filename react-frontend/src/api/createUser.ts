import AuthorType from "../models/author";

const endpoint = "http://localhost:4000";

async function postUser(userObject: AuthorType) {
  const json = JSON.stringify(userObject);

  console.log(userObject);

  const response = await fetch(`${endpoint}/users`, {
    method: "POST",
    body: json,
    headers: {
      "Content-Type": "application/json",
    },
  });

  return response;
}

export { postUser };
