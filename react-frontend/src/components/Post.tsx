import { useState, useEffect } from "react";
import { NavLink } from "react-router-dom";
import { PostType } from "../models/post";
import { AuthorType } from "../models/author";
import { getAuthor, getHasLiked, getLikeCount } from "../api/get";
import { likePost, unlikePost } from "../api/likePost";
import { deletePost } from "../api/deletePost";

interface PostProps {
  post: PostType;
  loggedIn: boolean;
}

export default function Post({ post, loggedIn }: PostProps) {
  const [thisPost, setThisPost] = useState<PostType>(post);
  const [author, setAuthor] = useState<AuthorType>({} as AuthorType);
  const [isPostDropdown, setIsPostDropdown] = useState(false);
  const [hasLiked, setHasLiked] = useState(false);
  const [likeCount, setLikeCount] = useState(0);

  useEffect(() => {
    findAuthor();
    findLikeCount();
    if (loggedIn) {
      getHasLiked(post).then((data) => {
        setHasLiked(data.hasLiked);
      });
    }
  }, []);

  function findLikeCount() {
    getLikeCount(post).then((data) => {
      setLikeCount(data.likes);
    });
  }

  function findAuthor() {
    getAuthor(post.authorId).then((data) => {
      setAuthor(data);
    });
  }

  const handleLikePost = () => {
    if (!loggedIn) {
      alert("Please log in to like this post.");
      return;
    }

    if (hasLiked) {
      unlikePost(thisPost)
        .then((response) => {
          setLikeCount(likeCount - 1);
          console.log("Post unliked successfully:", response);
          setHasLiked(false);
        })
        .catch((error) => {
          console.error("Error unliking post:", error);
        });
    } else {
      likePost(thisPost)
        .then((response) => {
          setLikeCount(likeCount + 1);
          console.log("Post liked successfully:", response);
          setHasLiked(true);
        })
        .catch((error) => {
          console.error("Error liking post:", error);
        });
    }
  };

  const dropdownClass = "absolute left-[52vw] z-10 mt-[-30px] w-48 rounded-md bg-white py-1 shadow-lg ring-1 ring-black ring-opacity-5 focus:outline-none transition ease-out duration-100";
  const closeDropdownClass = dropdownClass + "transform opacity-0 scale-95";
  const openDropdownClass = dropdownClass + "transform opacity-100 scale-100";

  const toggleDropdown = () => {
    setIsPostDropdown(!isPostDropdown);
  };

  return (
    <div className="bg-white p-8 rounded-lg shadow-md max-w-md mb-10">
      {/* <!-- User Info with Three-Dot Menu --> */}
      <div className="flex items-center justify-between mb-4">
        <div className="flex items-center space-x-2">
          <NavLink to={`/profile/${author.id}`}>
            <img src={author.avatar} alt="User Avatar" className="w-10 h-10 rounded-full" />
          </NavLink>

          <div>
            <NavLink to={`/profile/${author.id}`}>
              <p className="text-gray-800 font-semibold">{author.username}</p>
            </NavLink>
            <p className="text-gray-500 text-sm">
              Posted
              {new Date(thisPost.postDate).toLocaleString("en-GB", {
                dateStyle: "full",
                timeStyle: "long",
              })}
            </p>
          </div>
        </div>
        <div className="text-gray-500 cursor-pointer">
          {/* <!-- Three-dot menu icon --> */}
          <button onClick={toggleDropdown} className="hover:bg-gray-50 rounded-full p-1">
            <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" strokeWidth="2" strokeLinejoin="round">
              <circle cx="12" cy="7" r="1" />
              <circle cx="12" cy="12" r="1" />
              <circle cx="12" cy="17" r="1" />
            </svg>
          </button>
        </div>
      </div>
      <div className={isPostDropdown ? openDropdownClass : closeDropdownClass} role="menu" aria-orientation="vertical" aria-labelledby="user-menu-button" tabIndex={-1}>
        {/* <!-- Active: "bg-gray-100", Not Active: "" --> */}
        <a href="#" className="block px-4 py-2 text-sm text-gray-700 hover:bg-gray-100" role="menuitem" tabIndex={-1} id="user-menu-item-0">
          Edit Post
        </a>
        <a
          onClick={() => {
            deletePost(post.id);
            // UpdateGrid();
          }}
          href="#"
          className="block px-4 py-2 text-sm text-gray-700 hover:bg-gray-100"
          role="menuitem"
          tabIndex={-1}
          id="user-menu-item-1"
        >
          Delete Post
        </a>
      </div>
      {/* <!-- Message --> */}
      <div className="mb-4">
        <p className="text-gray-800">
          {thisPost.text}
          <a href="" className="text-blue-600">
            #CuteKitten
          </a>
          <a href="" className="text-blue-600">
            #AdventureCat
          </a>
        </p>
      </div>
      {/* <!-- Image --> */}
      <div className="mb-4">
        <img src={thisPost.imgSrc} alt="Post Image" className="w-full h-96 object-cover rounded-md" />
      </div>
      {/* <!-- Like and Comment Section --> */}
      <div className="flex items-center justify-between text-gray-500">
        <div className="flex items-center space-x-2">
          <button onClick={handleLikePost} className="flex justify-center items-center gap-2 px-2 hover:bg-gray-50 rounded-full p-1">
            <svg className={hasLiked ? "w-5 h-5 fill-red-900" : "w-5 h-5 fill-current"} xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24">
              <path d="M12 21.35l-1.45-1.32C6.11 15.36 2 12.28 2 8.5 2 5.42 4.42 3 7.5 3c1.74 0 3.41.81 4.5 2.09C13.09 3.81 14.76 3 16.5 3 19.58 3 22 5.42 22 8.5c0 3.78-4.11 6.86-8.55 11.54L12 21.35z" />
            </svg>
            <span>{likeCount} Likes</span>
          </button>
        </div>
        <button className="flex justify-center items-center gap-2 px-2 hover:bg-gray-50 rounded-full p-1">
          <svg width="22px" height="22px" viewBox="0 0 24 24" className="w-5 h-5 fill-current" xmlns="http://www.w3.org/2000/svg">
            <g id="SVGRepo_bgCarrier" strokeWidth="0"></g>
            <g id="SVGRepo_tracerCarrier" strokeLinejoin="round"></g>
            <g id="SVGRepo_iconCarrier">
              <path
                fillRule="evenodd"
                clipRule="evenodd"
                d="M12 22C17.5228 22 22 17.5228 22 12C22 6.47715 17.5228 2 12 2C6.47715 2 2 6.47715 2 12C2 13.5997 2.37562 15.1116 3.04346 16.4525C3.22094 16.8088 3.28001 17.2161 3.17712 17.6006L2.58151 19.8267C2.32295 20.793 3.20701 21.677 4.17335 21.4185L6.39939 20.8229C6.78393 20.72 7.19121 20.7791 7.54753 20.9565C8.88837 21.6244 10.4003 22 12 22ZM8 13.25C7.58579 13.25 7.25 13.5858 7.25 14C7.25 14.4142 7.58579 14.75 8 14.75H13.5C13.9142 14.75 14.25 14.4142 14.25 14C14.25 13.5858 13.9142 13.25 13.5 13.25H8ZM7.25 10.5C7.25 10.0858 7.58579 9.75 8 9.75H16C16.4142 9.75 16.75 10.0858 16.75 10.5C16.75 10.9142 16.4142 11.25 16 11.25H8C7.58579 11.25 7.25 10.9142 7.25 10.5Z"
              ></path>
            </g>
          </svg>
          <span>3 Comments</span>
        </button>
      </div>
    </div>
  );
}
