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

  const handleLoggedHours = () => {
    if (!roll) return;

    fetch("http://localhost:8080/log", {
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
