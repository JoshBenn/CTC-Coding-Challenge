"use client";

import Header from "./components/header"
import Content from "./components/content";
import { useState } from "react";
import { UserData } from "./models/user";

export default function Home() {
    const [userData, updateUserData] = useState<UserData | undefined>(undefined);

    return (
        <main className="flex h-screen flex-col items-center">
            <Header userData={userData} />
            <Content userData={userData} updateUserData={updateUserData} />
        </main>
    );
}
