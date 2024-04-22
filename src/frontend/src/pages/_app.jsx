import "@/styles/globals.css";
import { Raleway } from 'next/font/google';

const raleway = Raleway({
  subsets: ['latin'],
  weight: ['400', '700'],
  variable: '--font-raleway',
});

const MyApp = ({ Component, pageProps }) => {
  return (
    <main className={`${raleway.variable} font-sans`}>
      <Component {...pageProps} />
    </main>
  );
};

export default MyApp;
