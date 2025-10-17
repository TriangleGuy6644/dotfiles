import os
import discord
from discord import app_commands
from google.genai import Client
from dotenv import load_dotenv

load_dotenv()

DISCORD_TOKEN = os.getenv("DISCORD_TOKEN")
GUILD_ID = int(os.getenv("GUILD_ID"))
PROJECT_ID = os.getenv("GCP_PROJECT_ID")

intents = discord.Intents.default()
intents.message_content = True
client = discord.Client(intents=intents)
tree = app_commands.CommandTree(client)

genai_client = Client(project=PROJECT_ID)

@tree.command(
    name="summarize",
    description="Summarize the last [number] messages",
    guild=discord.Object(id=GUILD_ID)
)
@app_commands.describe(number="Number of messages to summarize (max 100)")
async def summarize(interaction: discord.Interaction, number: int = 100):
    if number > 100:
        number = 100
    channel = interaction.channel
    messages = [msg async for msg in channel.history(limit=number)]
    messages_text = "\n".join(msg.content for msg in reversed(messages) if msg.author != client.user)

    if not messages_text:
        await interaction.response.send_message("No messages to summarize")
        return

    response = genai_client.models.generate_content(
        model="gemini-2.5-flash",
        contents=f"Summarize the following Discord messages:\n{messages_text}"
    )

    summary = response.text
    await interaction.response.send_message(f"**Summary:**\n{summary}")

@client.event
async def on_ready():
    await tree.sync(guild=discord.Object(id=GUILD_ID))
    print(f"Logged in as {client.user}")

client.run(DISCORD_TOKEN)
