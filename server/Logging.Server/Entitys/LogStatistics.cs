﻿using MongoDB.Bson;

namespace Logging.Server
{
    public class LogStatistics
    {
        public ObjectId _id { get; set; }

        public int AppId { get; set; }

        public string AppName { get; set; }

        public int Debug { get; set; }

        public int Info { get; set; }

        public int Warn { get; set; }

        public int Error { get; set; }

        public long Time { get; set; }
    }
}