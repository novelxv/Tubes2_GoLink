import "@/styles/globals.css";
import type { AppProps } from "next/app";
import { Raleway } from 'next/font/google';

const raleway = Raleway({
  subsets: ['latin'],
  weight: ['400', '700'],
  variable: '--font-raleway',
});

function MyApp({ Component, pageProps }: AppProps) {
  return (
    <main className={`${raleway.variable} font-sans`}>
      <Component {...pageProps} />
    </main>
  )
}

export default MyApp;
