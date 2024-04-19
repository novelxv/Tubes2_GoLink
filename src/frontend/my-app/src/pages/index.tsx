import React from "react";
import Input from "@/components/entry";
import Entry from "@/components/entry";

export default function Home() {
  return (
    <main
      className={`bg-neutral-800 flex min-h-screen flex-col items-center justify-between p-24`}
    >
      <Entry/>
    </main>
  );
}
