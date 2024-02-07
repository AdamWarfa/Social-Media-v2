import NbaGameType from "../models/nbaGame";

interface NbaGameProps {
  game: NbaGameType;
}

export default function NbaGame({game}: NbaGameProps) {
  
  return (

    <div className="bg-white-950">
        <div className="flex justify-between p-4">
            <div className="flex">
            <div className="flex flex-col justify-center">
                <p className="text-2xl font-bold">{game.home_team.name}</p>
                <p className="text-sm">{game.home_team_score}</p>
            </div>
            <div className="flex flex-col justify-center ml-4">
                <p className="text-2xl font-bold">{game.visitor_team.name}</p>
                <p className="text-sm">{game.visitor_team_score}</p>
            </div>
            </div>
            <div className="flex flex-col justify-center">
            <p className="text-sm">{game.time}</p>
            <p className="text-sm">{game.date}</p>
            </div>
        </div>
    </div>

  );
}
