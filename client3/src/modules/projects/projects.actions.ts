import api from "../../api";
import projectsSlice, { type Project } from "./projects.slice";

export async function loadProjects() {
  try {
    const projects = await api.records.getFullList('projects', 200, {
      sort: '-created',
    }) as unknown as Project[] | null;

    if (!projects) {
      return;
    }

    (projects || []).forEach(projectsSlice.actions.set);

  } catch (error) {
    // TODO(bogdan): handle this
    console.log(error);
    console.error("ERROR loading projects");
  }
}
