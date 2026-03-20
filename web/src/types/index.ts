import z from "zod";

export type User = {
  id: string;
  email: string;
  username: string;
  displayname: string;
  imageUrl: string | null;
  color: string;
};

export type Profile = Omit<User, "email">;

export type Post = Definition & {
  id: string;
  visibility: Visibility;
  createdAt: string;
  updatedAt: string;
  likeCount: number;
  bookmarkCount: number;
  isLiked: boolean;
  isBookmarked: boolean;

  userId: string;
  username: string;
  displayname: string;
  imageUrl: string | null;
  color: string;
};

export const VISIBILITY = ["public", "private"] as const;
export type Visibility = (typeof VISIBILITY)[number];

export const PARTS_OF_SPEECH = ["noun", "verb", "adjective", "adverb"] as const;
export type PartOfSpeech = (typeof PARTS_OF_SPEECH)[number];

export type Definition = {
  word: string;
  definition: string;
  partOfSpeech: PartOfSpeech;
  pronunciation: string;
  synonyms: string[];
  antonyms: string[];
  example: string;
};

export const definitionSchemaOptional = z.object({
  word: z.string().optional(),
  pronunciation: z.string().optional(),
  definition: z.string().optional(),
  partOfSpeech: z.enum(PARTS_OF_SPEECH).optional(),
  synonyms: z.string().array().optional(),
  antonyms: z.string().array().optional(),
  example: z.string().optional(),
});

export const definitionSchema = z.object({
  word: z.string(),
  pronunciation: z.string(),
  definition: z.string(),
  partOfSpeech: z.enum(PARTS_OF_SPEECH),
  synonyms: z.string().array(),
  antonyms: z.string().array(),
  example: z.string(),
});
