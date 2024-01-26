import { useParams } from "react-router-dom";
import ProfileGrid from "../components/ProfileGrid";
import Nav from "../components/Nav";

interface ProfilePageProps {
  loggedIn: boolean;
  setLoggedIn: (value: boolean) => void;
  userId: string;
  setUserId: (value: string) => void;
}

export default function ProfilePage({ loggedIn, setLoggedIn, setUserId, userId }: ProfilePageProps) {
  const { id } = useParams();

  return (
    <>
      <Nav loggedIn={loggedIn} setLoggedIn={setLoggedIn} setUserId={setUserId} userId={userId} currentPage="signup/profile" />
      <h1>Profile Page</h1>
      <ProfileGrid authorId={id} />
    </>
  );
}
