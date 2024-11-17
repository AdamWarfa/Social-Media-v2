import { useParams } from "react-router-dom";
import ProfileGrid from "../components/ProfileGrid";

export default function ProfilePage() {
  const { id } = useParams();

  return (
    <>
      <h1>Profile Page</h1>
      <ProfileGrid authorId={id} />
    </>
  );
}
