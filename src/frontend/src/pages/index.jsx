import React from "react";
import Entry from "@/components/entry";
import Navbar from "@/components/navbar";
import Graph from "@/components/graph";
import Test from "@/components/test";

export default function Home() {
  return (
    <main className={`bg-neutral-800 flex flex-col items-center justify-between p-10`} style={{ minHeight: '100vh' }}>
      <Navbar/>
      <Entry />
      <Graph />
      <Test />
    </main>
  );
}