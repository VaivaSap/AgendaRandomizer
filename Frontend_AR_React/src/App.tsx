import { useEffect, useState } from "react";
import "./App.css";
import type { Activity } from "./Activity";

function App() {
  const [activities, setActivities] = useState<Activity[]>([]);
  const [roll, setRoll] = useState<Activity | null>(null);

  useEffect(() => {
    fetch("http://localhost:8080/activities")
      .then((response) => response.json())
      .then((data) => setActivities(data.activities))
      .catch((error) => console.error("Error fetching activities:", error));
  }, []);

  const handleRoll = () => {
    fetch("http://localhost:8080/randomize")
      .then((response) => response.json())
      .then((data) => setRoll(data))
      .catch((error) => console.error("Error rolling:", error));
  };

  return (
    <>
      <h1>Activities</h1>
      <ul>
        {activities.map((activity) => (
          <li key={activity.activity_name}>
            {activity.activity_name} || Target: {activity.target_hours} / Spent:{" "}
            {activity.logged_hours}
          </li>
        ))}
      </ul>

      <button onClick={handleRoll}>What to do next?</button>
      {roll && (
        <div>
          <h2>Rolled Activity</h2>
          <p>
            {roll.activity_name} || Target: {roll.target_hours} / Spent:{" "}
            {roll.logged_hours}
          </p>
        </div>
      )}
    </>
  );
}

export default App;
