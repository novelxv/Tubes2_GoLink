import React, {useState,useEffect} from "react";
import { BounceLoader } from "react-spinners";

const LoadingDots = () => {
    const [dots, setDots] = useState("");
  
    useEffect(() => {
      const intervalId = setInterval(() => {
        setDots((prevDots) => (prevDots.length >= 5 ? "" : prevDots + "."));
      }, 500);
  
      return () => clearInterval(intervalId);
    }, []);
  
    return <span className="font-raleway text-yellow-300 text-lg font-semibold animate-pulse">loading{dots}</span>;
  };
const Loading = () => {
    return (
        <div className="flex flex-col justify-center items-center p-5">
            <div className="flex flex-col justify-center items-center">
                <BounceLoader
                    color="#f6cc6e"
                    loading={true}
                    size={100}
                    className="opacity-0.25"
                />
                <span className="font-raleway text-yellow-300 text-lg font-semibold animate-pulse">We are calculating the results</span>
                <LoadingDots/>
            </div>
        </div>
        
    );
};

export default Loading;