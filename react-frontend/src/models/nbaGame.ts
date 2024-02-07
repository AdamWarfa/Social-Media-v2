import NbaTeamType from './nbaTeam';

interface NbaGameType {
    id: string;
    date: string;
    home_team: NbaTeamType;
    home_team_score: number;
    period: number;
    postseason: boolean;
    season: number;
    status: string;
    time: string;
    visitor_team: NbaTeamType;
    visitor_team_score: number;

}
export default NbaGameType;