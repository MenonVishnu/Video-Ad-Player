import "./VideoPlayer.css";
import Video from "../../Assets/video-1.mp4";
import { useState, useRef, useEffect } from "react";

function VideoPlayer(props) {
	const [ad, setAd] = useState(props.adsData[0]);
	const videoRef = useRef(null); // Reference to the video element

	useEffect(() => {
		const adInterval = setInterval(() => {
			setAd(props.adsData[      Math.floor( (Math.random() * 100) % props.adsData.length    )]);
		}, 10000); // Change ad every 10 seconds

		return () => clearInterval(adInterval);
	}, []);

	const handleAdClick = (event) => {
		// event.preventDefault();
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
				className="ad-overlay"
				onClick={handleAdClick}
				href={ad.target_url}
				target="_blank"
                rel="noreferrer"
                >
				<img src={ad.image_url} alt="" />
			</a>
		</div>
	);
}

export default VideoPlayer;

//next steps:
/*
    how to change positions randomly
*/
