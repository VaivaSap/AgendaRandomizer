import { useEffect, useState } from 'react'
import './App.css'
import type { Activity } from './Activity';

function App() {
  const [activities, setActivities] = useState<Activity[]>([]);

  useEffect(() => {
    fetch('http://localhost:8080/activities')
      .then(response => response.json())
      .then(data => setActivities(data.activities))
      .catch(error => console.error('Error fetching activities:', error));
  }, []);

  return (
    <>
      <h1>Activities</h1>
      <ul>
        {activities.map(activity => (
          <li key={activity.activity_name}>{activity.activity_name} || Target: {activity.target_hours} / Spent: {activity.logged_hours}</li>
        ))}
      </ul>
    </>
  )
}

export default App
