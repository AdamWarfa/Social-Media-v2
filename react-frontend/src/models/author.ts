interface AuthorType {
  id: string;
  avatar: string;
  followers: number;
  email: string;
  username: string;
  password: string;
}
interface UserRequest {
  email: string;
  username: string;
  password: string;
}

export type { AuthorType, UserRequest };
