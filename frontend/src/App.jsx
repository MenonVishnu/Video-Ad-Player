import { useEffect, useState } from "react";
import "./App.css";
import Header from "./components/Header/Header";
import VideoPlayer from "./components/VideoPlayer/VideoPlayer";

function App() {
	const [data, setData] = useState(null);

	useEffect(() => {
		async function fetchData() {
			try {
				//fetch advertisment
				const url = process.env.BACKEND_API || "localhost:8080";
				var response = await fetch(`http://${url}/api/v1/ads`);
				response = await response.json();
				setData(response.data);
			} catch (error) {
				console.log("API Connection failed!!", error);
			}
		}
		fetchData();
	}, []);

	return (
		<div className="App">
			<Header />
			{/* conditional rendering, only shows video if advData is loaded */}
			{data == null ? <h2>Loading...</h2> : <VideoPlayer advData={data} />}
		</div>
	);
}

export default App;
