import HomeGrid from "../components/HomeGrid";

interface HomePageProps {
  userId: string;
  setUserId: (value: string) => void;
}

export default function HomePage({ userId }: HomePageProps) {
  return (
    <>
      <HomeGrid userId={userId} />
    </>
  );
}
