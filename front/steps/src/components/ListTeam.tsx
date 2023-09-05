import { useState } from "react";

interface Props {
  teams: Team[];
  heading: string;
  onSelectTeam: (team: string) => void;
}
interface Team {
  id: string;
  name: string;
  counter: number;
  members: string[];
}
function ListTeam({ teams, heading, onSelectTeam }: Props) {
  const [selectedIndex, setSelectedIndex] = useState(-1);

  return (
    <>
      <h1>{heading}</h1>
      <ul className="list-group">
        {teams.map((team, index) => (
          <li
            className={
              selectedIndex === index
                ? "list-group-item active"
                : "list-group-item"
            }
            key={team.id}
            onClick={() => {
              setSelectedIndex(index);
              onSelectTeam(team.id);
            }}
          >
            {team.name}
          </li>
        ))}
      </ul>
    </>
  );
}

export default ListTeam;
