import Header from "../components/Header";
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
      <Header loggedIn={loggedIn} setLoggedIn={setLoggedIn} setUserId={setUserId} userId={userId} />
      <h1>Home Page</h1>
      <HomeGrid userId={userId} loggedIn={loggedIn} />
    </>
  );
}
