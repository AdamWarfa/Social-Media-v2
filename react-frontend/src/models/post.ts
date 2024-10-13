interface PostType {
  id: string;
  authorId: string;
  text: string;
  imgSrc: string;
  postDate: string;
  likes: number;
}

interface PostRequest {
  authorId: string;
  text: string;
  imgSrc: string;
}

export type { PostType, PostRequest };
