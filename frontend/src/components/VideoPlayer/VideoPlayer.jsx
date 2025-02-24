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
			setPosition("img-" + randIndex);
		}, 10000);

		return () => clearInterval(advInterval);
	}, []);

	const handleAdvClick = async () => {
		if (videoRef.current) {
			try {
				const timestamp = videoRef.current.currentTime;

				//sending click data
				const response = await fetch("http://localhost:8080/api/v1/ads/click", {
					method: "POST",
					body: JSON.stringify({
						ad_id: adv.ad_id,
						timestamp: String(new Date().getTime()),
						ip: "",
						video_timestamp: timestamp,
					}),
				});
			} catch (error) {
				console.log("Error sending click data: ", error);
			}
		}
	};

	return (
		<div className="container">
			{/* Video Player  */}
			<video ref={videoRef} className="video" controls>
				<source src={Video} type="video/mp4" />
			</video>

			{/* Ad Overlay  */}
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

