import HomeGrid from "../components/HomeGrid";

interface HomePageProps {
  loggedIn: boolean;
  setLoggedIn: (value: boolean) => void;
  userId: string;
  setUserId: (value: string) => void;
}

export default function HomePage({ loggedIn, userId }: HomePageProps) {
  return (
    <>
      <HomeGrid userId={userId} loggedIn={loggedIn} />
    </>
  );
}
