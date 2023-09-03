// import { useState, useEffect } from "react";
// import useAxiosPrivate from "../hooks/useAxiosPrivate";
// import { useNavigate, useLocation } from "react-router-dom";

const Users = () => {
  // const [users, setUsers] = useState();
  // const axiosPrivate = useAxiosPrivate();
  // const navigate = useNavigate();
  // const location = useLocation();

  // useEffect(() => {
  //     let isMounted = true;
  //     const controller = new AbortController();

  //     const getUsers = async () => {
  //         try {
  //             const response = await axiosPrivate.get('/users', {
  //                 signal: controller.signal
  //             });
  //             console.log(response.data);
  //             isMounted && setUsers(response.data);
  //         } catch (err) {
  //             console.error(err);
  //             navigate('/login', { state: { from: location }, replace: true });
  //         }
  //     }

  //     getUsers();

  //     return () => {
  //         isMounted = false;
  //         controller.abort();
  //     }
  // }, [])
  const users = [
    {
      name: "eaylmer0",
      company: "Meezzy",
      role: "Financial Analyst",
      is_verified: true,
      status: "BANNED",
    },
    {
      name: "amerveille1",
      company: "Thoughtstorm",
      role: "Senior Sales Associate",
      is_verified: false,
      status: "BANNED",
    },
    {
      name: "iocollopy2",
      company: "Katz",
      role: "Quality Control Specialist",
      is_verified: true,
      status: "ACTIVE",
    },
    {
      name: "jeuler3",
      company: "Mydo",
      role: "Community Outreach Specialist",
      is_verified: false,
      status: "BANNED",
    },
    {
      name: "ktothe4",
      company: "Mydo",
      role: "Web Designer I",
      is_verified: false,
      status: "BANNED",
    },
    {
      name: "hcoard5",
      company: "Topicshots",
      role: "Nuclear Power Engineer",
      is_verified: false,
      status: "",
    },
    {
      name: "klorek6",
      company: "Skyble",
      role: "Junior Executive",
      is_verified: false,
      status: "ACTIVE",
    },
    {
      name: "rbier7",
      company: "Quatz",
      role: "Nurse",
      is_verified: false,
      status: "",
    },
    {
      name: "fkoopman8",
      company: "Agivu",
      role: "Nuclear Power Engineer",
      is_verified: false,
      status: "BANNED",
    },
    {
      name: "heustes9",
      company: "Eadel",
      role: "Statistician III",
      is_verified: true,
      status: "ACTIVE",
    },
    {
      name: "rspinksa",
      company: "Babblestorm",
      role: "Engineer III",
      is_verified: true,
      status: "BANNED",
    },
    {
      name: "abourrelb",
      company: "Gabspot",
      role: "Marketing Assistant",
      is_verified: true,
      status: "",
    },
    {
      name: "fianielloc",
      company: "Feedfire",
      role: "Structural Analysis Engineer",
      is_verified: false,
      status: "ACTIVE",
    },
    {
      name: "vbraxtond",
      company: "Bubbletube",
      role: "VP Sales",
      is_verified: false,
      status: "",
    },
    {
      name: "amarke",
      company: "Gabtype",
      role: "Design Engineer",
      is_verified: false,
      status: "BANNED",
    },
    {
      name: "lpagetf",
      company: "Yozio",
      role: "Web Developer I",
      is_verified: false,
      status: "BANNED",
    },
    {
      name: "vhorbartg",
      company: "InnoZ",
      role: "Analog Circuit Design manager",
      is_verified: true,
      status: "BANNED",
    },
    {
      name: "rcollardh",
      company: "Tekfly",
      role: "Chief Design Engineer",
      is_verified: true,
      status: "",
    },
    {
      name: "zdarwenti",
      company: "Flipopia",
      role: "GIS Technical Architect",
      is_verified: false,
      status: "",
    },
    {
      name: "sgauchej",
      company: "Cogidoo",
      role: "Administrative Officer",
      is_verified: true,
      status: "ACTIVE",
    },
  ];

  console.log(users[1]);
  return (
    <>
      <article>
        <h2>Users List</h2>
        {users?.length ? (
          <ul>
            {users.map((user, i) => (
              <li key={i}>{user?.name}</li>
            ))}
          </ul>
        ) : (
          <p>No users to display</p>
        )}
      </article>
    </>
  );
};

export default Users;
