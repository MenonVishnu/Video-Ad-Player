import "./VideoPlayer.css";
import Video from "../../Assets/video-1.mp4";

function VideoPlayer() {
	return (
		<div className="container">
			<video className="video" controls>
				<source src={Video} type="video/mp4" />
				Your Browser does not support this video
			</video>

            <div className="ad-overlay">
                <h2>This is ad display</h2>
            </div>
		</div>
	);
}

export default VideoPlayer;

