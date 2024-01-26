import { useParams } from "react-router-dom";
import Header from "../components/Header";
import ProfileGrid from "../components/ProfileGrid";

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
      <Header loggedIn={loggedIn} setLoggedIn={setLoggedIn} setUserId={setUserId} userId={userId} />
      <h1>Profile Page</h1>
      <ProfileGrid authorId={id} />
    </>
  );
}
