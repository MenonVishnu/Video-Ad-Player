import "./VideoPlayer.css";
import Video from "../../Assets/video-1.mp4";
import { useState, useRef, useEffect } from "react";

function VideoPlayer(props) {
	const [adv, setAdv] = useState(props.advData[0]);
	const videoRef = useRef(null); // Reference to the video element

	useEffect(() => {

        //Change Ad every 10 seconds
		const advInterval = setInterval(() => {
			setAdv(props.advData[      Math.floor( (Math.random() * 100) % props.advData.length    )]);
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
				className="adv-overlay"
				onClick={handleAdvClick}
				href={adv.target_url}
				target="_blank"
                rel="noreferrer"
                >
				<img src={adv.image_url} alt="" />
			</a>
		</div>
	);
}

export default VideoPlayer;

//next steps:
/*
    How to show ads random order for 10 sec / or loop them continueously
    how to change positions randomly
*/
