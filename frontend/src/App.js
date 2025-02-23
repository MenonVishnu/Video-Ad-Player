import { useEffect, useState } from "react";
import "./App.css";
import Header from "./components/Header/Header";
import VideoPlayer from "./components/VideoPlayer/VideoPlayer";

function App() {
	const [data, setData] = useState(null);

	useEffect(() => {
		async function fetchData() {
			try {
				//todo
				var response = await fetch("http://localhost:8080/api/v1/ads");
				response = await response.json();
				console.log(await response.data);
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
			{data == null ? <h2>Loading</h2> : <VideoPlayer advData={data} /> }
		</div>
	);
}

export default App;
