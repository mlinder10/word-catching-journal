import type { Definition } from "@/types";
import { useNavigate } from "@tanstack/react-router";
import { Card, CardContent, CardFooter } from "@/components/ui/card";
import { ArrowRight } from "lucide-react";
import DefinitionView from "@/components/definition-view";
import { Button } from "@/components/ui/button";

type DefinitionItemProps = {
  definition: Definition;
};

export default function DefinitionItem({ definition }: DefinitionItemProps) {
  const navigate = useNavigate();

  function handlePost() {
    navigate({
      to: "/app/post",
      search: definition,
    });
  }

  return (
    <li>
      <Card className="flex-col h-full">
        <CardContent className="flex-1 space-y-2">
          <DefinitionView definition={definition} />
        </CardContent>
        <CardFooter className="py-2">
          <Button onClick={handlePost}>
            <span>Post</span>
            <ArrowRight />
          </Button>
        </CardFooter>
      </Card>
    </li>
  );
}
