import { useEffect, useState } from "react";
import "./App.css";
import Header from "./components/Header/Header";
import VideoPlayer from "./components/VideoPlayer/VideoPlayer";

function App() {
	const [data, setData] = useState(null);

	useEffect(() => {
		async function fetchData() {
			try {
				const response = await fetch("http://localhost:8080/api/v1/ads");
				setData(response.json());
			} catch (error) {
				console.log("API Connection failed!!", error);
			}
		}
		fetchData();
	}, []);

	// const data = [
	// 	{
	// 		ad_id: 1,
	// 		image_url:
	// 			"https://media.istockphoto.com/id/1408387701/photo/social-media-marketing-digitally-generated-image-engagement.jpg?s=612x612&w=0&k=20&c=VVAxxwhrZZ7amcPYJr08LLZJTyoBVMN6gyzDk-4CXos=",
	// 		target_url: "https://www.google.com",
	// 	},
	// 	{
	// 		ad_id: 2,
	// 		image_url:
	// 			"https://images.pexels.com/photos/276267/pexels-photo-276267.jpeg?auto=compress&cs=tinysrgb&w=1260&h=750&dpr=2",
	// 		target_url: "https://www.facebook.com",
	// 	},
	// 	{
	// 		ad_id: 3,
	// 		image_url:
	// 			"https://images.pexels.com/photos/3965543/pexels-photo-3965543.jpeg?auto=compress&cs=tinysrgb&w=1260&h=750&dpr=2",
	// 		target_url: "https://www.github.com",
	// 	},
	// 	{
	// 		ad_id: 4,
	// 		image_url:
	// 			"https://images.pexels.com/photos/3965543/pexels-photo-3965543.jpeg?auto=compress&cs=tinysrgb&w=1260&h=750&dpr=2",
	// 		target_url: "https://www.reddit.com",
	// 	},
	// ];

	return (
		<div className="App">
			<Header />
			{/* conditional rendering, only shows video if advData is loaded */}
			{data ? <VideoPlayer advData={data} /> : <></>}
		</div>
	);
}

export default App;
