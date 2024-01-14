import { useRef, useState } from "react";

export default function App() {
  return (
    <>
      <GetUser />
      <CreateUser />
    </>
  );
}

function GetUser() {
  const id = useRef<HTMLInputElement>(null);
  const [user, setUser] = useState<
    { username: string; id: number } | undefined
  >(undefined);
  const [messages, setMessages] = useState<string[]>([]);
  return (
    <>
      <input ref={id} />
      <input
        type="submit"
        value="Get User"
        onClick={() =>
          fetch(`/api/user/${id.current?.value}`)
            .then((res) => res.json())
            .then((data) => setUser(data))
            .catch((err) =>
              setMessages([
                ...messages,
                `Error getting user ${id.current?.value}: ${JSON.stringify(err)}`,
              ])
            )
        }
      />
      {user && <div>{JSON.stringify(user)}</div>}
      <ul>
        {messages.map((message, i) => (
          <li key={i}>{message}</li>
        ))}
      </ul>
    </>
  );
}

function CreateUser() {
  const [username, setUsername] = useState("");
  const [messages, setMessages] = useState<string[]>([]);
  return (
    <>
      <input value={username} onChange={(e) => setUsername(e.target.value)} />
      <input
        type="submit"
        value="Create User"
        onClick={() =>
          fetch("/api/user", {
            method: "POST",
            headers: { "Content-Type": "application/json" },
            body: JSON.stringify({ username }),
          })
            .then((res) => res.json())
            .then((data) =>
              setMessages([
                ...messages,
                `Success created user ${username}: ${JSON.stringify(data)}`,
              ])
            )
            .catch((err) =>
              setMessages([
                ...messages,
                `Error creating user ${username}: ${JSON.stringify(err)}`,
              ])
            )
        }
      />
      <ul>
        {messages.map((message, i) => (
          <li key={i}>{message}</li>
        ))}
      </ul>
    </>
  );
}
