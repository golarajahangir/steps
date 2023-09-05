import ListTeam from "./components/ListTeam";

function App() {
  let teams = [
    { id: "1", name: "Team 1", counter: 0, members: ["Alice", "Bob", "Carol"] },
    { id: "2", name: "Team 2", counter: 0, members: ["Teo", "Tony"] },
    { id: "3", name: "Team 3", counter: 0, members: ["Alexander", "Sara"] },
  ];
  const handleSelectTeam = (team: string) => {
    console.log(team);
  };
  return (
    <ListTeam
      teams={teams}
      heading="Teams List"
      onSelectTeam={handleSelectTeam}
    />
  );
}
export default App;
