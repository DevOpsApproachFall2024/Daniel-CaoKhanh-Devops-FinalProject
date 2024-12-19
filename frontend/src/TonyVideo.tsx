import React from 'react';
import './ytvideos.css';

const TonyVideo : React.FC = () => {
     return (
         <div className="video-container">
             <h1>Embedded YouTube Video</h1>
             <iframe width="560" height="315" src="https://www.youtube.com/embed/aAkMkVFwAoo?si=1gcu44ZcUckuSuh6"
                     title="YouTube video player" frameBorder="0"
                     allow="accelerometer; autoplay; clipboard-write; encrypted-media; gyroscope; picture-in-picture; web-share"
                     allowFullScreen></iframe>
         </div>
     );
};

export default TonyVideo;

