from __future__ import annotations

from typing import Dict, List, Literal, Optional, TypedDict

from langchain_core.messages import BaseMessage, HumanMessage
from langgraph.graph import END, StateGraph


class AgentState(TypedDict):
    messages: List[BaseMessage]
    financial_context: Dict[str, float]
    eligibility_status: Optional[bool]
    proposal_draft: str


def _mock_llm_route(state: AgentState) -> Literal["Researcher", "Compliance", "Writer", "End"]:
    if not state.get("financial_context"):
        return "Researcher"
    if state.get("eligibility_status") is None:
        return "Compliance"
    if state.get("eligibility_status") is True and not state.get("proposal_draft"):
        return "Writer"
    return "End"


def supervisor_node(state: AgentState) -> Literal["Researcher", "Compliance", "Writer", "End"]:
    return _mock_llm_route(state)


def researcher_node(state: AgentState) -> AgentState:
    financial_context = dict(state.get("financial_context", {}))
    financial_context.update({"total_revenue": 500000, "net_assets": 100000})
    state["financial_context"] = financial_context

    messages = list(state.get("messages", []))
    messages.append(HumanMessage(content="Data gathered."))
    state["messages"] = messages
    return state


def compliance_node(state: AgentState) -> AgentState:
    net_assets = state.get("financial_context", {}).get("net_assets", 0)
    state["eligibility_status"] = net_assets > 0
    return state


def writer_node(state: AgentState) -> AgentState:
    net_assets = state.get("financial_context", {}).get("net_assets", 0)
    state["proposal_draft"] = (
        f"Draft Proposal based on ${net_assets:,.0f} net assets and the organization's financial profile."
    )
    return state


def build_agentic_graph() -> StateGraph:
    graph = StateGraph(AgentState)
    graph.add_node("Supervisor", supervisor_node)
    graph.add_node("Researcher", researcher_node)
    graph.add_node("Compliance", compliance_node)
    graph.add_node("Writer", writer_node)

    graph.add_conditional_edges(
        "Supervisor",
        supervisor_node,
        {
            "Researcher": "Researcher",
            "Compliance": "Compliance",
            "Writer": "Writer",
            "End": END,
        },
    )

    graph.set_entry_point("Supervisor")
    graph.set_finish_point("Writer")
    return graph
