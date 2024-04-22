import React, {useState,useEffect} from "react";
import { BounceLoader } from "react-spinners";

const LoadingDots: React.FC = () => {
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
            <div className="absolute top-3/4 left-1/3 w-[571px] h-[442px] bg-rose-400 rounded-full blur-[200px]  -z-1"></div>
            <div className="absolute top-3/4 left-1/2 w-[571px] h-[442px] bg-emerald-400 rounded-full blur-[200px]  -z-1"></div>
            <div className="flex flex-col justify-center items-center z-0">
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
