import "./VideoPlayer.css";
import Video from "../../Assets/video-1.mp4";
import { useState, useRef, useEffect } from "react";

function VideoPlayer(props) {
	const [adv, setAdv] = useState(props.advData[0]);
	const [position, setPosition] = useState("img-0");
	const videoRef = useRef(null); // Reference to the video element

	useEffect(() => {
		//Change Ad every 10 seconds
		const advInterval = setInterval(() => {
			//change the adv randmoly
			const randIndex = Math.floor(Math.random() * props.advData.length);
			setAdv(props.advData[randIndex]);

			//change position of the element
			const position = randIndex % 4;
			setPosition("img-" + position);
		}, 10000);

		return () => clearInterval(advInterval);
	}, []);

	const handleAdvClick = () => {
		if (videoRef.current) {
			const timestamp = videoRef.current.currentTime;
			console.log("Ad clicked at timestamp:", timestamp);
		}
	};

	return (
		<div className="container">
			<video ref={videoRef} className="video" controls>
				<source src={Video} type="video/mp4" />
				Your Browser does not support this video
			</video>

			<a
				id="adv-container"
				className="adv-overlay"
				onClick={handleAdvClick}
				href={adv.target_url}
				target="_blank"
				rel="noreferrer">
				<img src={adv.image_url} alt="" className={position} />
			</a>
		</div>
	);
}

export default VideoPlayer;

//next steps:
/*
	Integration with backend
	Axios: connect to backend api and retrieve ads and send click data

*/
