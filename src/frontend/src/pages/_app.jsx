import "@/styles/globals.css";
import { Raleway } from 'next/font/google';
import { Toaster } from "@/components/ui/toaster";
import ParticlesComponent from "@/components/particles";

const raleway = Raleway({
  subsets: ['latin'],
  weight: ['400', '700'],
  variable: '--font-raleway',
});

const MyApp = ({ Component, pageProps }) => {
  return (
    <main className={`${raleway.variable} font-sans`}>
      <ParticlesComponent/>
      <Toaster/>
      <Component {...pageProps} />
    </main>
  );
};

export default MyApp;
