import { useEffect, useState } from "react";
import "./App.css";
import type { Activity } from "./Activity";

const API = `http://${window.location.hostname}:8080`;

function App() {
  const [activities, setActivities] = useState<Activity[]>([]);
  const [roll, setRoll] = useState<Activity | null>(null);

  useEffect(() => {
    fetch(`${API}/activities`)
      .then((response) => response.json())
      .then((data) => setActivities(data.activities))
      .catch((error) => console.error("Error fetching activities:", error));
  }, []);

  const handleRoll = () => {
    fetch(`${API}/randomize`)
      .then((response) => response.json())
      .then((data) => setRoll(data))
      .catch((error) => console.error("Error rolling:", error));
  };

  const handleLoggedHours = () => {
    if (!roll) return;

    fetch(`${API}/log`, {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify({
        activity_name: roll.activity_name,
        hours_spent: 1,
      }),
    })
      .then((response) => {
        if (!response.ok) {
          throw new Error("Log failed: " + response.status);
        }

        setActivities(
          activities.map((activity) =>
            activity.activity_name === roll.activity_name
              ? { ...activity, logged_hours: activity.logged_hours + 1 }
              : activity,
          ),
        );

        setRoll({ ...roll, logged_hours: roll.logged_hours + 1 });
      })
      .catch((error) => console.error("Error logging hours:", error));
  };

  return (
    <>
      <h1>Activities</h1>
      <ul>
        {activities.map((activity) => (
          <li key={activity.activity_name}>
            {activity.activity_name} || Target: {activity.target_hours} / Spent:{" "}
            {activity.logged_hours}
            <progress
              value={activity.logged_hours}
              max={activity.target_hours}
            ></progress>
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
          <button onClick={handleLoggedHours}>
            Done 1h: {roll.activity_name}
          </button>
        </div>
      )}
    </>
  );
}

export default App;
