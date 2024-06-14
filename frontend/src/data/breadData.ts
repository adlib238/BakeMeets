import { Bread } from "../types/Bread";

const breadData: Bread[] = [
  {
    id: 1,
    name: "バケット",
    imageUrl: "img/bread1.webp",
    description: "食感がハードで、強い小麦の香りがする",
    likes: 120,
    shares: 45,
    comments: 20,
  },
  {
    id: 2,
    name: "ベーグル",
    imageUrl: "img/bread2.webp",
    description: "もっちりとしていて、ほんのり甘い生地とゴマがマッチしている",
    likes: 200,
    shares: 60,
    comments: 30,
  },
  {
    id: 3,
    name: "パンオショコラ",
    imageUrl: "img/bread3.webp",
    description: "パリパリと音が鳴るような生地と中の甘いチョコレートが美味しい",
    likes: 180,
    shares: 55,
    comments: 25,
  },
];

export default breadData;
