const endpoint: string = "http://localhost:4000";

async function getNbaGame() {
    const date = new Date();
    const year = date.getFullYear();
    const month = (date.getMonth() + 1).toString().padStart(2, '0');
    const day = date.getDate().toString().padStart(2, '0');

    const formattedDate = `${year}-${month}-${day}`;


    const res = await fetch(`${endpoint}/nba/games/${formattedDate}`);
    const data = await res.json();
    return data;
}


export { getNbaGame };