import { API_URL } from "../settings.ts";
import { handleHttpErrors } from "./fetchUtils.ts";
const LOGIN_URL = API_URL + "/users/login";

export type User = { id: string; username: string; password: string; roles?: string[] };

interface LoginResponse {
  username: string;
  id: string;
  token: string;
  roles: Array<string>;
}

interface LoginRequest {
  username: string;
  password: string;
}

const authProvider = {
  isAuthenticated: false,
  signIn(user_: LoginRequest): Promise<LoginResponse> {
    const options: RequestInit = {
      method: "POST",
      headers: {
        "Content-type": "application/json",
        Accept: "application/json",
      },
      body: JSON.stringify(user_),
    };
    return fetch(LOGIN_URL, options).then(handleHttpErrors);
  },
};

export type { LoginResponse, LoginRequest };
export { authProvider };
