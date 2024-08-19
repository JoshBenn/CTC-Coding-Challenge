"use client";

import Header from "./components/header"
import Content from "./components/content";
import { useState } from "react";

export default function Home() {
    const [userData, updateUserData] = useState(null);
    return (
        <main className="flex h-screen flex-col items-center">
            <Header userData={userData} />
            <Content userData={userData} />
        </main>
    );
}
