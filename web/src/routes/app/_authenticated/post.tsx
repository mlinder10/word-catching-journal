import PageTitle from "@/components/custom-ui/page-title";
import DefinitionForm from "@/components/pages/post/definition-form";
import PostForm from "@/components/pages/post/post-form";
import { definitionSchema, definitionSchemaOptional } from "@/types";
import { createFileRoute } from "@tanstack/react-router";
import z from "zod";

export const Route = createFileRoute("/app/_authenticated/post")({
  validateSearch: definitionSchemaOptional,
  component: PostPage,
});

function parseDefinition(data: z.infer<typeof definitionSchemaOptional>) {
  const result = definitionSchema.safeParse(data);
  if (result.success) return result.data;
  return null;
}

function PostPage() {
  const search = Route.useSearch();
  const definition = parseDefinition(search);

  return (
    <main className="flex flex-col h-full">
      <PageTitle title="Post" />
      {definition ? <PostForm definition={definition} /> : <DefinitionForm />}
    </main>
  );
}
