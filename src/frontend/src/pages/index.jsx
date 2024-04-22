import React from "react";
import Entry from "@/components/entry";

export default function Home() {
  return (
    <main className={`bg-neutral-800 flex flex-col items-center justify-between p-24`} style={{ minHeight: '100vh' }}>
      <Entry />
    </main>
  );
}