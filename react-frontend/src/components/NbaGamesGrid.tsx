import { useState ,useEffect } from "react";
import { getNbaGame } from "../api/getNba";
import NbaGameType from "../models/nbaGame";
import NbaGame from "./NbaGame";

export default function NbaGamesGrid() {
    const [games, setGames] = useState<NbaGameType[]>([]);

  useEffect(() => {
    function fecthTodaysGames() {
      getNbaGame().then((data) => {
        setGames(data);
      });
    }
    fecthTodaysGames();
  }, []);

  
  return (

    <div className="mt-20">
        {games.map((game: NbaGameType) => (
            <NbaGame key={game.id} game={game}/>
        ))}
    </div>
  );
}
