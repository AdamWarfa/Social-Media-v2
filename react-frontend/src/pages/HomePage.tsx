import Nav from "../components/Nav";
import HomeGrid from "../components/HomeGrid";

interface HomePageProps {
  loggedIn: boolean;
  setLoggedIn: (value: boolean) => void;
  userId: string;
  setUserId: (value: string) => void;
}

export default function HomePage({ loggedIn, setLoggedIn, userId, setUserId }: HomePageProps) {
  return (
    <>
      <Nav loggedIn={loggedIn} setLoggedIn={setLoggedIn} setUserId={setUserId} userId={userId} currentPage="homepage" />
      <HomeGrid userId={userId} loggedIn={loggedIn} />
    </>
  );
}
